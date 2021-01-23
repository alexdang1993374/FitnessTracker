import React, { Component } from "react";
import axios from "axios";

const User = (props) => (
  <tr>
    <td>{props.user.username}</td>
    <td>
      <a
        href="#"
        onClick={() => {
          props.deleteUser(props.user.id);
        }}
      >
        delete
      </a>
    </td>
  </tr>
);

export default class CreateUser extends Component {
  constructor(props) {
    super(props);

    this.onChangeUsername = this.onChangeUsername.bind(this);
    this.onSubmit = this.onSubmit.bind(this);
    this.deleteUser = this.deleteUser.bind(this);

    this.state = {
      users: [],
      username: "",
    };
  }

  componentDidMount() {
    axios
      .get("http://localhost:5000/api/users")
      .then((res) => {
        console.log(res.data.data);
        if (res.data.data !== null) {
          this.setState({
            users: res.data.data,
          });
        }
      })
      .catch((error) => {
        console.log(error);
      });
  }

  onChangeUsername(e) {
    this.setState({
      username: e.target.value,
    });
  }

  onSubmit(e) {
    e.preventDefault();

    const user = {
      username: this.state.username,
    };

    console.log(user);

    axios
      .post("http://localhost:5000/api/users", user)
      .then((res) => console.log(res.data));

    setTimeout(() => {
      axios.get("http://localhost:5000/api/users").then((res) => {
        console.log(res.data.data);
        if (res.data.data.length > 0) {
          this.setState({
            users: res.data.data,
          });
        }
      });
    }, 500);

    this.setState({
      username: "",
    });
  }

  deleteUser(id) {
    axios
      .delete("http://localhost:5000/api/users/" + id)
      .then((res) => console.log(res.data));

    this.setState({
      users: this.state.users.filter((el) => el.id !== id),
    });
  }

  userList() {
    return this.state.users.map((currentUser) => {
      return (
        <User
          user={currentUser}
          deleteUser={this.deleteUser}
          key={currentUser.id}
        />
      );
    });
  }

  render() {
    return (
      <>
        <h3>Create New User</h3>
        <form onSubmit={this.onSubmit}>
          <div className="form-group">
            <label>Username: </label>
            <input
              type="text"
              required
              className="form-control"
              value={this.state.username}
              onChange={this.onChangeUsername}
            />
          </div>
          <div className="form-group">
            <input
              type="submit"
              value="Create user"
              className="btn btn-primary"
            />
          </div>
        </form>
        <h3>Current Users</h3>
        <table className="table">
          <thead className="thead-light">
            <tr>
              <th>Username</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>{this.userList()}</tbody>
        </table>
      </>
    );
  }
}
