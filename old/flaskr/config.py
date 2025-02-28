from flask import Flask
from flask_sqlalchemy import SQLAlchemy

# creating the application itself
app = Flask(
    __name__,
    template_folder="templates",
    static_folder="static",
    static_url_path="/static",
)
app.secret_key = "my_secret_key_1234567890"

# specify which databases we will use
app.config["SQLALCHEMY_DATABASE_URI"] = "sqlite:///database.db"
app.config["SQLALCHEMY_BINDS"] = {
    "clients": "sqlite:///clients.db",
    "rentals": "sqlite:///rentals.db",
}

db = SQLAlchemy(app)
