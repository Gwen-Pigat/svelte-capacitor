package initializers

import (
	"database/sql"
	"os"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("DB_URI"))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	err = SetupDB(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func SetupDB(db *sql.DB) error {
	dbInit := `SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";
CREATE TABLE IF NOT EXISTS ` + "`task`" + ` (
  ` + "`id`" + ` int NOT NULL AUTO_INCREMENT,
  ` + "`date_add`" + ` datetime NOT NULL,
  ` + "`date_to`" + ` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  ` + "`title`" + ` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  ` + "`content`" + ` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  ` + "`is_done`" + ` tinyint(1) NOT NULL,
  ` + "`ref_user`" + ` int DEFAULT NULL,
  PRIMARY KEY (` + "`id`" + `),
  KEY ` + "`ref_user`" + ` (` + "`ref_user`" + `)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
CREATE TABLE IF NOT EXISTS ` + "`user`" + ` (
  ` + "`id`" + ` int NOT NULL AUTO_INCREMENT,
  ` + "`username`" + ` varchar(255) NOT NULL,
  ` + "`date_add`" + ` datetime DEFAULT NULL,
  ` + "`is_active`" + ` tinyint(1) NOT NULL,
  ` + "`token`" + ` varchar(255) NOT NULL,
  PRIMARY KEY (` + "`id`" + `)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;`
	_, err := db.Exec(dbInit)
	if err != nil {
		return err
	}
	return nil
}
