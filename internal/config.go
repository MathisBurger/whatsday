package internal

import (
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types"
	"gopkg.in/yaml.v2"
	"os"
)

type FullConfig struct {
	Birthdays map[string][]string `yaml:"birthdays"`
}

func (c *FullConfig) GetConfig() *FullConfig {

	yamlFile, err := os.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		panic(err)
	}

	return c
}

func (c *FullConfig) GetJIDMap(client *whatsmeow.Client) map[string][]types.JID {
	idMap := make(map[string][]types.JID)
	for k, v := range c.Birthdays {
		resp, err := client.IsOnWhatsApp(v)
		if err != nil {
			panic(err)
		}
		var ids []types.JID
		for _, jid := range resp {
			ids = append(ids, jid.JID)
		}
		idMap[k] = ids
	}
	return idMap
}
