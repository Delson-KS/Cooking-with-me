package tests

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/tebeka/selenium"
)

//java -jar selenium-server-4.29.0-SNAPSHOT.jar standalone

const (
	seleniumPath  = "chromedriver.exe"                  // Путь к chromedriver
	port          = 9516                                // Порт Selenium Standalone Server
	loginURL      = "http://localhost:8080/loginPage"   // URL страницы логина
	recipePageURL = "http://localhost:8080/recipesPage" // URL страницы Pokémon
)

func TestRecipePage(t *testing.T) {
	// Запуск ChromeDriver локально
	service, err := selenium.NewChromeDriverService("/opt/homebrew/bin/chromedriver", port) // Укажите путь к chromedriver
	if err != nil {
		log.Fatalf("Ошибка запуска ChromeDriver: %v", err)
	}
	defer service.Stop()

	// Подключение к локальному ChromeDriver
	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, "http://localhost:9516")
	if err != nil {
		t.Fatalf("Ошибка подключения к WebDriver: %v", err)
	}
	defer wd.Quit()

	// Открываем страницу логина
	err = wd.Get("http://localhost:8080/loginPage")
	if err != nil {
		t.Fatalf("Ошибка перехода на страницу логина: %v", err)
	}

	fmt.Println("Тест успешно запустился")
}

// Функция ожидания появления элемента с таймаутом
func waitForElement(wd selenium.WebDriver, selector string, timeout time.Duration) error {
	// Ожидание до timeout
	var err error
	endTime := time.Now().Add(timeout)
	for time.Now().Before(endTime) {
		_, err = wd.FindElement(selenium.ByCSSSelector, selector)
		if err == nil {
			return nil
		}
		time.Sleep(500 * time.Millisecond) // Ожидание перед следующим запросом
	}
	return err
}
