from flask import render_template, request, redirect, url_for
from config import app, db
from models import VhsTape


with app.app_context():
    db.create_all()


@app.route("/")
def home():
    return render_template("home_page.html")


@app.route("/vhs/<int:id>")
def vhs(id):
    vhs = VhsTape.query.get(id)
    return render_template("vhs_page.html", vhs=vhs)


@app.route("/all_vhstapes")
def all_vhstapes():
    all_vhstapes = VhsTape.query.order_by(VhsTape.year.desc()).all()
    return render_template("all_vhstape_page.html", all_vhstapes=all_vhstapes)


@app.route("/create_vhstape", methods=["GET", "POST"])
def create_vhstape():
    if request.method == "POST":
        name_film = request.form["name_film"]
        year = request.form["year"]
        age_rating = request.form["age_rating"]
        count = request.form["count"]

        vhs = VhsTape(
            name_film=name_film, year=year, age_rating=age_rating, count=count
        )

        try:
            db.session.add(vhs)
            db.session.commit()
            return redirect("/create_vhstape")
        except Exception as err:
            return f"There was an issue adding your VHS tape:/n {str(err)}"
    else:
        return render_template("create_vhstape_page.html")


@app.route("/vhs/<int:id>/update", methods=["GET", "POST"])
def update(id):
    vhs = VhsTape.query.get(id)
    if request.method == "POST":
        name_film = request.form["name_film"]
        year = request.form["year"]
        age_rating = request.form["age_rating"]
        count = request.form["count"]

        if name_film:
            vhs.name_film = name_film
        if year:
            vhs.year = year
        if age_rating:
            vhs.age_rating = age_rating
        if count:
            vhs.count = count

        try:
            db.session.commit()
            return redirect(url_for("vhs", id=id))
        except Exception as err:
            return f"failed to update:/n {str(err)}"

    else:
        return render_template("update_vhstape_page.html", vhs=vhs)


@app.route("/delete_vhstape/<int:id>")
def delete_vhstape(id):
    vhs = VhsTape.query.get_or_404(id)

    try:
        db.session.delete(vhs)
        db.session.commit()
        return redirect("/all_vhstapes")
    except Exception as err:
        return f"There was a problem deleting that VHS tape:/n {str(err)}"

@app.route("/clear_database", methods=["POST"])
def clear_database():
    db.session.query(VhsTape).delete()
    db.session.commit()

    return "The database has been cleared"


if __name__ == "__main__":
    app.run(debug=True)
