# vk-subscribe-checker

Консольное приложение проверяет список пользователей на возможности написать им сообщение через сообщения сообщества
[vk.com/dev/messages.isMessagesFromGroupAllowed](https://vk.com/dev/messages.isMessagesFromGroupAllowed)

### Usage

- У вас должен быть файл с id пользователей в ВК, каждый id с новой строки или через запятую.
- Потребуется API ключ сообщества с правами message. Получить его можно в Управлении сообществом > Работа с API
- Скачайте необходимый файл тут [github.com/stels-cs/vk-subscribe-checker/releases/tag/1.0.0](https://github.com/stels-cs/vk-subscribe-checker/releases/tag/1.0.0)

Пример (testIds.txt – файл с id пользователей 0775a1581c – API ключ сообщества)

```bash
./vk-subscribe-checker-darwin-amd64 --input testIds.txt  --token 0775a1581c > allow.txt
2018/02/16 02:55:22 Checking ids for group TestGroup
2018/02/16 02:55:22 DONE
2018/02/16 02:55:22 Total ids in file 30 #30 id было в файле
2018/02/16 02:55:22 Ids who's allow write 29 #29 id группа может написать, сами id в файле allow
````
