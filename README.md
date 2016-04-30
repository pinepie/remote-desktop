# remote-desktop
## Naming
* host - устройство к которому осуществляется доступ
* client - устройство с которого производится доступ к хосту
* server - соединяет клиент и хост 

## Actions
Описание видов действий, и формата данных в сообщениях. Для простоты опускается упоминание полей action и data.
Если не описывается структура поля data то это строка. Сгрупированы по получателю\обработчику.

**Server:**
* HOST_REGISTER - Регистрация хоста на сервере. Хост отправляет свой логин.
* GET_HOSTS - Запрос клиентом списка доступных хостов у сервера.
* SELECT_HOST - Выбор клиентом хоста для соединения. Передается имя хоста
* CLIENT_ACCESS - Ответ на CLIENT_CONNECT. Хост разрешает соединение с клиентом. Сервер соединяет хост с клиентом.
* CLIENT_DENIED - Ответ на CLIENT_CONNECT. Хост запрещает соединение с клиентом. Сервер продолжает обработку сообщений хоста.

**Host:**
* MOUSE_MOVE - Перемещение мыши.
* MOUSE_RPRESS - Зажать правую кнопку мыши.
* MOUSE_RRELEASE - Отпустить правую кнопку мыши.
* MOUSE_LPRESS - Зажать левую кнопку мыши.
* MOUSE_LRELEASE - Отпустить левую кнопку мыши.
* CLIENT_CONNECT - Клиент хочет установить подключение с хостом. Хост должен ответить CLIENT_ACCESS или CLIENT_DENIED
* CLIENT_CLOSE - Сервер сообщает хосту что клиент закрыл соединение.

**Client:**
* IMG_FRAME - Кадр с хоста.
* SELECT_SUCCESS - Ответ на SELECT_HOST. Успешное соединение с хостом.
* HOST_BUSY - Ответ на SELECT_HOST. Выбраный хост в данный момент занят.
* AVALIABLE_HOSTS - Ответ на GET_HOSTS. Передается масив имен доступных хостов.
