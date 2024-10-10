# app file

from flaskr import app, db
from flaskr import route_home_and_others, route_vhs, route_client, route_rental

# creating databases, if databases already exist, they will not be created
with app.app_context():
    db.create_all()

# registering routers
app.register_blueprint(route_home_and_others)
app.register_blueprint(route_vhs)
app.register_blueprint(route_client)
app.register_blueprint(route_rental)

if __name__ == "__main__":
    app.run(debug=False, host='0.0.0.0')
