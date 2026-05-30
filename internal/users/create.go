package users

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"text/template"

	"github.com/sanikum3/ORP/internal/configs"
	"github.com/sanikum3/ORP/internal/models"
	"github.com/sanikum3/ORP/internal/storage"
	"github.com/sanikum3/ORP/internal/systemd"
	"github.com/sanikum3/ORP/internal/templates"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

const (
	BaseDir = "/opt/olcrtc-users"
)

var JitsiBase = "meet.handyweb.org"

type Config struct {
	RoomID string
	Key    string
}

func RandomString(length int) (string, error) {
	bytes := make([]byte, length)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	for i := range bytes {
		bytes[i] = charset[int(bytes[i])%len(charset)]
	}

	return string(bytes), nil
}

func RandomHex(bytesLength int) (string, error) {
	bytes := make([]byte, bytesLength)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}

func CreateUser(name string) error {
	roomid, _ := RandomString(16)
	fmt.Println("User created:", name, "roomid: ", roomid)
	var hash, err = RandomHex(32)
	if err != nil {
		fmt.Println("Error generating hash:", err)
		return err
	}
	cfg, err := configs.Load()
	if err != nil {
		return err
	}

	fullroom := fmt.Sprintf(
		"https://%s/%s",
		cfg.JitsiBase,
		roomid,
	)
	userdir := fmt.Sprintf("%s/%s", BaseDir, name)
	fmt.Println("Full room URL:", fullroom)
	config := Config{
		RoomID: roomid,
		Key:    hash,
	}
	err = os.Mkdir(userdir, 0755)
	if err != nil {
		fmt.Println("Error creating user directory:", err)
		return err
	}
	content, err := templates.FS.ReadFile("srv_template.yaml")
	if err != nil {
		return fmt.Errorf("read template: %w", err)
	}
	template, err := template.New("config").Parse(string(content))
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return err
	}
	file, err := os.Create(userdir + "/srv_template.yaml")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()
	err = template.Execute(file, config)
	if err != nil {
		return fmt.Errorf("template: %w", err)
	}
	url := fmt.Sprintf(
		"olcrtc://jitsi?datachannel@%s#%s%%%s$OlcRTC",
		fullroom,
		hash,
		name,
	)
	err = systemd.Create(name)
	if err != nil {
		return fmt.Errorf("create systemd: %w", err)
	}

	err = systemd.Reload()
	if err != nil {
		return fmt.Errorf("reload systemd: %w", err)
	}

	err = systemd.Enable(name)
	if err != nil {
		return fmt.Errorf("enable service: %w", err)
	}

	err = systemd.Restart(name)
	if err != nil {
		return fmt.Errorf("restart service: %w", err)
	}
	storage.Users[name] = models.User{
		Username: name,
		Room:     fullroom,
		Hash:     hash,
		Url:      url,
	}
	err = storage.Save()
	if err != nil {
		return fmt.Errorf("save %w", err)
	}
	fmt.Println("UserCreated", name)
	fmt.Println()
	fmt.Println("URL:")
	fmt.Println(url)

	return nil
}
