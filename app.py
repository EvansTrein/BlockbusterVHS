from flask import Flask, render_template, request, redirect
from flask_sqlalchemy import SQLAlchemy


app = Flask(__name__)

app.config["SQLALCHEMY_DATABASE_URI"] = "sqlite:///database.db"

db = SQLAlchemy(app)


class VhsTape(db.Model):
    id_num = db.Column(db.Integer, primary_key=True)
    name_film = db.Column(db.String(100), nullable=False)
    year = db.Column(db.Integer, nullable=False)
    age_rating = db.Column(db.String(3), nullable=False)
    count = db.Column(db.Integer, nullable=False)

with app.app_context():
    db.create_all()


@app.route("/")
def home():
    return render_template("home_page.html")

@app.route("/vhs/<int:id>")
def vhs(id):
    return render_template("vhs_page.html")

@app.route("/create_vhstape", methods=["GET", "POST"])
def create_vhstape():
    if request.method == "POST":
        name_film = request.form["name_film"]
        year = request.form["year"]
        age_rating = request.form["age_rating"]
        count = request.form["count"]

        vhs = VhsTape(name_film=name_film, 
                    year=year, 
                    age_rating=age_rating, 
                    count=count)
        
        db.session.add(vhs)
        db.session.commit()
        return redirect("/create_vhstape")
    else:
        return render_template("create_vhstape_page.html")


if __name__ == "__main__":
    app.run(debug=True)