## FOX Storage

- Fast
- Online
- eXtendable

## これはなに？

GDGoC Japan Convention 2025 に向けて制作した，自作 NAS 的なもの．

バックエンドとして Go を採用し，Web のバックエンドっぽく構築した．

Raspberry Pi による低価格のクラウドストレージのようなものをやるために作成．

紹介スライド：https://docs.google.com/presentation/d/1DIV8lmpbftCeQhmYITYpxzwH53Er0P6AQ90MQvlwHHQ/edit?usp=sharing

## 技術スタック

- Go
  - gin
  - gorm
  - jwt
- Next.js

  - shadcn/ui
  - （ほとんど v0 が作ってくれた）

- Web サーバ，DB サーバ（mariadb）を docker-compose で立ち上げています．

## 注意事項

- `config/`に pem を生成する必要があります．
- 本リポジトリは現状有姿で公開され，何らの保証もありません．ご了承ください．
- `FOX`は適当に付けた名前なので「なにが`Fast`」やねん「なにが`eXtendable`」やねんといったところは許してください．

## LICENSE に関する表示

MIT License

Copyright (c) 2025 Shadcn IO
Copyright (c) 2023 shadcn

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
