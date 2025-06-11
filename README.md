# 🎮 solitaire-serve-api
## 📌 概要 / Overview
このプロジェクトは放置型収集ゲームのバックエンドAPIです。<br>
Goを用いて設計・実装しました。<br>
目的：Goの習熟のため

---

## 🛠 使用技術 / Tech Stack
- Go 1.24.4
- MySQL
- Redis

---

## ✨ 主な機能 / Features
- プレイヤーデータの管理（村・資源・建物・兵士）
- サーバー定期処理（資源増加・戦闘処理）
- 攻撃予約 → 結果通知（非同期戦闘）
- リーダーボード表示

---

## 🚀 セットアップ手順 / How to Run
```bash
git clone ...
cd solitaire-serve-api
go run main.go # サーバーの起動
```

## API
- プレイヤー作成

```bash
curl -X POST http://localhost:8080/player \
  -H "Content-Type: application/json" \
  -d '{"id":"user123", "name":"Taro", "resources":100, "soldiers":10, "village":"StarterVille"}'
```

- プレイヤー情報の取得

```bash
curl http://localhost:8080/player?id=user123 #作成したユーザーIDを指定
```

---

## 📘API仕様
（準備中）

---

## 🏗️アーキテクチャ図
（準備中）

---

## 👤 担当範囲 / My Role
- 全て担当

---

## 📝 今後の課題 / ToDo
- データベース（PostgreSQL, Redis）
- ジョブキューによる非同期戦闘処理
- JWTログイン認証
- WebSocketによる通知
- クラウドデプロイ（Docker + AWS）

---

## 📅 開発期間
2025年06月〜現在（継続開発中）

## 📫 連絡先 / Contact
- GitHub: [https://github.com/tomo1015](https://github.com/tomo1015)
- メール: [tomo_a0901@outlook.jp]

     
