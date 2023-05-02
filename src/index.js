const express = require("express");
const app = express();
const cors = require("cors");
const pool = require("./db");

app.use(cors());
app.use(express.json());

// send a new chat
app.post("/chats", async (req, res) => {
    try {
        const {description} = req.body;
        const newChat = await pool.query("INSERT INTO chat (description) VALUES($1) RETURNING *", [description]);
    } catch (err) {
        console.error(err.message);
    }
});

// view all chats
app.get("/chats", async (req, res) => {
    try {
        const allChats = await pool.query("SELECT * FROM chat");
        res.json(allChats.rows);
    } catch (err) {
        console.error(err.message);
    }
});

app.listen(5000, function () {
    console.log("server has started on port 5000");
});
