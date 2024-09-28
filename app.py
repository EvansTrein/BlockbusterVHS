from flask import render_template, request, redirect, flash
from config import app, db
from models import VhsTape


# with app.app_context():
#     db.create_all()


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


        # checking the conditions for a database record
        if VhsTape.query.filter_by(name_film=name_film).first():
            flash('There\'s already a tape like this')
            return render_template("create_vhstape_page.html")
        elif not all((name_film, year, age_rating, count)):
            flash('All fields must be filled in')
            return render_template("create_vhstape_page.html")
        elif len(name_film) > 100 or len(age_rating) > 4:
            flash('The number of characters in the "name film" or "age rating" field is exceeded')
            return render_template("create_vhstape_page.html")
        else:
            try:
                vhs = VhsTape(name_film=name_film, year=year, age_rating=age_rating, count=count)   
                db.session.add(vhs)
                db.session.commit()
                return redirect("/create_vhstape")
            except Exception as err:
                return f"There was an issue adding your VHS tape:/n {str(err)}"
            
    else:
        return render_template("create_vhstape_page.html")


@app.route("/vhs/<int:id>/update", methods=["GET", "POST"])
def update_vhstape(id):
    vhs = VhsTape.query.get(id)
    if request.method == "POST":
        name_film = request.form["name_film"]
        year = request.form["year"]
        age_rating = request.form["age_rating"]
        count = request.form["count"]

        if all((name_film, year, age_rating, count)):
            vhs.name_film = name_film
            vhs.year = year
            vhs.age_rating = age_rating
            vhs.count = count
        else:
            flash('fields cannot be empty')
            return redirect(f"/vhs/{id}/update")

        try:
            db.session.commit()
            return redirect(f"/vhs/{id}")
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
