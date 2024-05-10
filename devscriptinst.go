package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

//Validation d'une étape en appuyant sur entrée
func pressKeyToContinue() {
	fmt.Println("Appuyez sur 'Entrée' pour continuer...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}


//Téléchargement d'un fichier
func downloadFile(filepath string, url string) error {
    fmt.Printf("Démarrage du téléchargement pour %s depuis %s\n", filepath, url)
    pressKeyToContinue()


    out, err := os.Create(filepath)
    if err != nil {
        return fmt.Errorf("erreur lors de la création du fichier %s : %w", filepath, err)
    }
    defer out.Close()


    resp, err := http.Get(url)
    if err != nil {
        return fmt.Errorf("erreur lors de la requête HTTP pour %s : %w", url, err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("échec de la requête HTTP : statut %s", resp.Status)
    }
    written, err := io.Copy(out, resp.Body)
    if err != nil {
        return fmt.Errorf("erreur lors de l'écriture des données dans le fichier %s : %w", filepath, err)
    }

    if written == 0 {
        return fmt.Errorf("aucune donnée écrite dans le fichier %s, vérifiez l'URL et la connexion réseau", filepath)
    }

    fmt.Printf("Téléchargement réussi et fichier enregistré : %s\n", filepath)
    pressKeyToContinue()
    return nil
}


//Vérification de l'existence d'un fichier
func fileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		fmt.Printf("Erreur lors de la vérification de l'existence du fichier %s : %s\n", filepath, err)
		pressKeyToContinue() 
	} 
	return true
}


//Téléchargement du fichier si besoin
func downloadIfNeeded(filepath, url string) {
	if fileExists(filepath) {
		fmt.Printf("Le fichier %s existe déjà. Pas de telechargement.\n", filepath)
		pressKeyToContinue()
	} else {
		if err:= downloadFile(filepath, url); err != nil {
			fmt.Printf("Erreur lors du telechargement du fichier %s : %s\n", filepath, err)
			pressKeyToContinue()
		}
	}
}


//Source d'exécutables
func main() {
    downloads := map[string]string{
        "VSCodeInstaller.exe": "https://update.code.visualstudio.com/latest/win32-x64-user/stable",
        "nvm-setup.exe":       "https://github.com/coreybutler/nvm-windows/releases/download/1.1.12/nvm-setup.exe",
        "GoInstaller.msi":     "https://dl.google.com/go/go1.19.2.windows-amd64.msi",
        "PHP.zip":             "https://windows.php.net/downloads/releases/php-8.3.6-nts-Win32-vs16-x64.zip",
		"rustup-init.exe":     "https://static.rust-lang.org/rustup/dist/x86_64-pc-windows-msvc/rustup-init.exe",
		"Git-2.45.0_64bit.exe": "https://github.com/git-for-windows/git/releases/download/v2.45.0.windows.1/Git-2.45.0-64-bit.exe",
		"tabby-1.0.207-setup-x64.exe": "https://github.com/Eugeny/tabby/releases/download/v1.0.207/tabby-1.0.207-setup-x64.exe",
		"Python-3.12.3-amd64.exe": "https://www.python.org/ftp/python/3.12.3/python-3.12.3-amd64.exe",
		"Node-v20.12.2-x64.msi": "https://nodejs.org/dist/v20.12.2/node-v20.12.2-x64.msi",

    }

    fmt.Println("Début des téléchargements...")
    pressKeyToContinue()
    for filepath, url := range downloads {
        downloadIfNeeded(filepath, url)
    }
    fmt.Println("Tous les téléchargements sont terminés.")
}