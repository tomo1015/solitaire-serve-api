# 🎮 solitaire-serve-api
## 📌 概要 / Overview
このプロジェクトは、「Go」と「SQLite」で構築した非同期PvEゲームです<br>
プレイヤーは兵士を育成し、ワールドマップ上にある防衛ポイントに対して攻撃予約を行います<br>
その後、予約時に指定した兵士の情報に従って自動的かつ非同期で戦闘を行います<br>

---

## 🛠 使用技術 / Tech Stack
- Go 1.24.4
- SQLite
---

## 🔹 主な機能 / Features
### プレイヤーデータの管理（村・資源・建物・兵士）
- SQLiteのオートインクリメントに従ってユニークに採番された兵士レコード
- 訓練時にレベルに応じて能力値が変化
### 攻撃予約
- 指定されたユニットタイプ、ユニット数に従ってバトル用兵士の情報を作成（訓練中の情報が影響しないように）
- 攻撃予約を行うことて一定時間後に非同期解決を実施
### バトル
- 予約されたユニットの攻撃力・防御力・クリティカル確率・命中率に従って結果を算出
- 戦果ポイントやルートアイテムも付与

---

## ✨ アピールポイント / Appeal Point
- Goの並行性を活用した非同期の攻撃予約・解決機能
- JSONやSQLiteを用いることによるデータ管理
- 兵士のレベルや能力値に基づくダメージ算出に加え、「命中率」と「クリティカル」といった確率も盛り込み<br>
従来の単純な数値対決に深みをプラスしました
- 攻撃予約時に戦闘に用いるユニット数を指定可能にすることで、プレイヤーの戦術に幅を持たせています。
- SQLiteのオートインクリメントを活用してデータ管理も整理しました
- バトルを行うことで獲得できる資源を使用して、村をアップグレードすることができます

## 🚀 セットアップ手順 / How to Run
```bash
git clone ...
cd solitaire-serve-api
go run main.go # サーバーの起動
```

## 🏗️構成図
```bash
solitaire-serve-api/
├─game.db
├─go.mod
├─go.sum
├─main.go
├─README.md
├─data
│  ├─defense_point.json
│  ├─defense_resource.json
│  └─defense_soldier.json
├─docs
│  ├─docs.go
│  ├─swagger.json
│  └─swagger.yaml
├─internal
│  ├─battle
│  │ └─battle.go
│  ├─db
│  │ ├─db.go
│  │ └─game.db
│  ├─handlers
│  │ ├─battle_handler.go
│  │ ├─building_handler.go
│  │ ├─leaderboard_handler.go
│  │ ├─login_handler.go
│  │ └─player_handler.go
│  │ └─soldier_handler.go
│  ├─leaderboard
│  │ └─leaderboard.go
│  ├─models
│  │ ├─attack.go
│  │ ├─building.go
│  │ ├─defensePoint.go
│  │ ├─player.go
│  │ ├─session.go
│  │ ├─soldier.go
│  │ └─worldMap.go
│  ├─scheduler
│  │ └─tasks.go
│  └─util
│    └─resource.go
└─storage
   ├─init.go
   └─memory.go

```
---

## 📘API仕様
セットアップを行い、サーバーを起動した後に以下のリンクで確認することができます。

http://localhost:8080/swagger/index.html#/
---

## 👤 担当範囲 / My Role
- 全て担当

---

## 📝 今後の課題 / ToDo
- Redis
- JWTログイン認証
- WebSocketによる通知
- クラウドデプロイ（Docker + AWS）

---

## 📅 開発期間
2025年06月〜現在（継続開発中）

## 📫 連絡先 / Contact
- GitHub: [https://github.com/tomo1015](https://github.com/tomo1015)
- メール: [tomo_a0901@outlook.jp]

     
