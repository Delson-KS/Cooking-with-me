package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/tebeka/selenium"
)

const (
	seleniumPath  = "chromedriver"                      // Путь к ChromeDriver (без .exe для macOS/Linux)
	port          = 9516                                // Локальный порт ChromeDriver (не используется, так как WebDriver запущен отдельно)
	loginURL      = "http://localhost:8080/loginPage"   // URL страницы логина
	recipePageURL = "http://localhost:8080/recipesPage" // URL страницы рецептов
)

func TestRecipePage(t *testing.T) {
	// Подключение к Selenium Server (замена порта 9516 на 4444)
	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, "http://localhost:4444/wd/hub")
	if err != nil {
		t.Fatalf("Ошибка подключения к WebDriver: %v", err)
	}
	defer wd.Quit()

	// Открываем страницу логина
	err = wd.Get(loginURL)
	if err != nil {
		t.Fatalf("Ошибка перехода на страницу логина: %v", err)
	}

	fmt.Println("Тест успешно запустился")
}

// Функция ожидания появления элемента с таймаутом
func waitForElement(wd selenium.WebDriver, selector string, timeout time.Duration) error {
	var err error
	endTime := time.Now().Add(timeout)
	for time.Now().Before(endTime) {
		_, err = wd.FindElement(selenium.ByCSSSelector, selector)
		if err == nil {
			return nil
		}
		time.Sleep(500 * time.Millisecond)
	}
	return err
}
