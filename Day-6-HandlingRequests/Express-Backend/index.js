const express = require("express");
const app = express();
const PORT = 5000;

app.use(express.json());
app.use(express.urlencoded({extended: true}));

app.get(express.json());
app.use(express.urlencoded({extended: true}));

app.get("/",(req,res) => {
    res.status(200).send("Welcome to Backend");
});

app.get("/get",(req,res) => {
    res.status(200).json({message: "Hello from debadritaghosh.tech"});
});

app.post("/post",(req,res) => {
    let myJson = req.body;
    res.status(200).send(myJson);
});

app.post("/postform",(req,res) => {
    res.status(200).send(JSON.stringify(req.body));
});

app.listen(PORT,() => {
    console.log("Server is running!")
})