import React, { Component } from "react";
import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
import { CookiesProvider } from "react-cookie";
import { Provider } from "mobx-react";

import {
  Dashboard,
  LoginPage,
  QuestionPage,
  RegisterPage,
  SupportPage
} from "./pages";
import { AdminRouter } from "./routers";
import * as stores from "./stores";

class App extends Component {
  render() {
    return (
      <div>
        <Provider {...stores}>
          <CookiesProvider>
            <Router>
              <Switch>
                <Route exact path="/" render={() => <LoginPage />} />
                <Route exact path="/register" render={() => <RegisterPage />} />
                <Route exact path="/dashboard" render={() => <Dashboard />} />
                <Route
                  exact
                  path="/question"
                  render={props => <QuestionPage {...props} />}
                />
                <Route path="/___admin" component={AdminRouter} />
                <Route exact path="/support" render={() => <SupportPage />} />
              </Switch>
            </Router>
          </CookiesProvider>
        </Provider>
      </div>
    );
  }
}

export default App;
