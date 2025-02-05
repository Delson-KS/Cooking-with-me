package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/tebeka/selenium"
)

const (
	loginURL      = "http://localhost:8080/loginPage"   // URL страницы логина
)

func TestRecipePage(t *testing.T) {
	caps := selenium.Capabilities{"browserName": "chrome"}
	if err != nil {
		t.Fatalf("Ошибка подключения к WebDriver: %v", err)
	}
	defer wd.Quit()

	// Открываем страницу логина
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
	}
	return err
}
