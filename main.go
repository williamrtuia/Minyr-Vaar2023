package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "os"
    "strconv"
    "strings"

    "github.com/williamrtuia/funtemps"
)

func main() {
    for {
        fmt.Println("Velkommen til minyr!")
        fmt.Println("Skriv 'convert' for å konvertere temperaturer fra Celsius til Fahrenheit.")
        fmt.Println("Skriv 'average' for å beregne gjennomsnittstemperaturer.")
        fmt.Println("Skriv 'exit' for å avslutte programmet.")

        scanner := bufio.NewScanner(os.Stdin)
        scanner.Scan()
        input := scanner.Text()

        switch input {
        case "convert":
            // Les inn data fra filen
            data, err := ioutil.ReadFile("kjevik-temp-celsius-20220318-20230318.csv")
            if err != nil {
                fmt.Println("Kunne ikke lese inn filen. Feilmelding:", err)
                continue
            }

            // Konverter temperaturer og lagre i ny fil
            lines := strings.Split(string(data), "\n")
            var convertedLines []string
            for i, line := range lines {
                if i == 0 {
                    // Kopier header-linjen til ny fil
                    convertedLines = append(convertedLines, line)
                } else {
                    fields := strings.Split(line, ";")
                    celsius, err := strconv.ParseFloat(fields[3], 64)
                    if err != nil {
                        // Kunne ikke konvertere til float
                        continue
                    }
                    fahr := conv.CelsiusToFahrenheit(celsius)
                    fields[3] = fmt.Sprintf("%.2f", fahr)
                    convertedLine := strings.Join(fields, ";")
                    convertedLines = append(convertedLines, convertedLine)
                }
            }
            newFileData := strings.Join(convertedLines, "\n")
            err = ioutil.WriteFile("kjevik-fahr-celsius-20220318-20230318.csv", []byte(newFileData), 0644)
            if err != nil {
                fmt.Println("Kunne ikke skrive til ny fil. Feilmelding:", err)
                continue
            }
            fmt.Println("Konvertering ferdig. Ny fil lagret.")

        case "average":
            // Les inn data fra filen
            data, err := ioutil.ReadFile("kjevik-temp-celsius-20220318-20230318.csv")
            if err != nil {
                fmt.Println("Kunne ikke lese inn filen. Feilmelding:", err)
                continue
            }

              // Beregn gjennomsnittstemperaturer
            lines := strings.Split(string(data), "\n")
            var celsiusTemps []float64
            for i, line := range lines {
                if i == 0 {
                    // Hopp over header-linjen
                    continue
                } else {
                    fields := strings.Split(line, ";")
                    celsius, err := strconv.ParseFloat(fields[3], 64)
                    if err != nil {
                        // Kunne ikke konvertere til float
                        continue
                    }
                    celsiusTemps =                     append(celsiusTemps, celsius)
                }
            }
            avgTemp := conv.AverageTemperature(celsiusTemps)
            fmt.Printf("Gjennomsnittstemperaturen er %.2f grader Celsius.\n", avgTemp)

        case "exit":
            fmt.Println("Ha en fin dag!")
            return

        default:
            fmt.Println("Ugyldig input. Vennligst prøv igjen.")
        }
    }
}

