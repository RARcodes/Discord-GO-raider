package main

import (
    "bufio"
    "bytes"
    "fmt"
    "net/http"
    "os"
    "sync"
)

var client = &http.Client{}

func sendMessage(channelID string, message []byte, botToken string, wg *sync.WaitGroup) {
    defer wg.Done()

    req, err := http.NewRequest("POST", "https://discord.com/api/v10/channels/"+channelID+"/messages", bytes.NewBuffer(message))
    if err != nil {
        panic(err)
    }

    req.Header.Add("Authorization", botToken)
    req.Header.Add("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println(resp.Status)
}

func main() {
    channelID := "974276548645503006" //change for id you want
    message := []byte(`{"content": "join discord.gg/cheapcord"}`) // change for your message you want xx

    for {
        file, err := os.Open("tokens.txt")
        if err != nil {
            panic(err)
        }
        defer file.Close()

        var wg sync.WaitGroup

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            botToken := scanner.Text()

            wg.Add(1)
            go sendMessage(channelID, message, botToken, &wg)

        }
        if err := scanner.Err(); err != nil {
            panic(err)
        }

        wg.Wait()

    }