from flaskr import app, db
from flaskr import route_home_and_others, route_vhs, route_client, route_rental

with app.app_context():
    db.create_all()

app.register_blueprint(route_home_and_others)
app.register_blueprint(route_vhs)
app.register_blueprint(route_client)
app.register_blueprint(route_rental)

if __name__ == "__main__":
    app.run(debug=True)
