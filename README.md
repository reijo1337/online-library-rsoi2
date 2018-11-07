# online-library-rsoi2
[![Build Status](https://travis-ci.org/reijo1337/online-library-rsoi2.svg?branch=develop)](https://travis-ci.org/reijo1337/online-library-rsoi2)

## ПРИМЕРЫ ЗАПРОСОВ

## GET
> getUserArrears

Параметры: name, size (опционально), page (опционально)

Описание: Возвращает записанные на читателя книги, возможна паггинация

Возвращаемые коды: 200, 400, 404

## POST
> newArear

Параметры: reader, book

Описание: Создание новой записи на читателя

Возвращаемые коды: 200, 400, 404, 500

## DELETE
> closeArrear

Параметры: id

Описание: Закрывает запись читателя (возвращение книги в библиотеку)

Возвращаемые коды: 200, 400, 404, 500
