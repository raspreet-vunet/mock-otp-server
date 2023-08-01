package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type UserOTP struct {
	Username string `json:"username"`
	OTP      string `json:"otp"`
}

// loadOTPs loads OTPs from JSON files in the given directory.
func loadOTPs(dir string) (map[string]string, error) {
	userOTPs := make(map[string]string)

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %v", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if strings.HasSuffix(file.Name(), ".json") {
			data, err := os.ReadFile(filepath.Join(dir, file.Name()))
			if err != nil {
				return nil, fmt.Errorf("failed to read file: %v", err)
			}

			var otps []UserOTP
			if err := json.Unmarshal(data, &otps); err != nil {
				return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
			}

			for _, otp := range otps {
				userOTPs[otp.Username] = otp.OTP
			}
		}
	}

	return userOTPs, nil
}

func startServer(port string, userOTPs map[string]string) {
    http.HandleFunc("/otp", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            log.Println("Invalid request method received")
            http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
            return
        }

        // Decode the JSON request body into a UserOTP struct
        var otpData UserOTP
        if err := json.NewDecoder(r.Body).Decode(&otpData); err != nil {
            log.Println("Bad request received")
            http.Error(w, "Bad request", http.StatusBadRequest)
            return
        }

        // Check if the OTP is correct
        expectedOTP, ok := userOTPs[otpData.Username]
        if !ok || otpData.OTP != expectedOTP {
            log.Printf("Invalid OTP received for user: %s\n", otpData.Username)
            http.Error(w, "Invalid username or OTP", http.StatusUnauthorized)
            return
        }

        log.Printf("OTP verified for user: %s\n", otpData.Username)
        fmt.Fprintf(w, "OTP verified for user %s", otpData.Username)
    })

    log.Printf("Starting server on port %s", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}




func main() {
    log.SetOutput(os.Stdout)
    dataDir, httpPort := GetConfig()

    // Load the OTPs
    userOTPs, err := loadOTPs(dataDir)
    if err != nil {
        log.Fatalf("Failed to load OTPs: %v", err)
    }
    log.Println("Successfully loaded OTPs")

    // Start the HTTP server
    startServer(httpPort, userOTPs)
}


