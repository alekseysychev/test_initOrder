Тест для очередности запуска `init()` пакетов при импорте

Результат:

```main -> mainsubPackage1 -> subPackage -> init()
main -> mainSubPackage1 -> init()
main -> mainSubPackage2 -> init()
main -> init()
main -> main()
main -> main() -> mainSubPackage2 -> Show()
```

![картинка](https://github.com/alekseysychev/test_initOrder/blob/main/image.jpg?raw=true)

[developpaper.com](https://developpaper.com/detailed-explanation-of-init-function-in-go-language/)

[golang.org](https://golang.org/doc/effective_go#init)
