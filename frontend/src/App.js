import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import Navigation from './components/Navigation';
import Dashboard from './pages/Dashboard';
import GatheringDetail from './pages/GatheringDetail';
import Login from './pages/Login';
import Register from './pages/Register';

const App = () => {
  return (
    <Router>
      <div className="app">
        <Navigation />
        <Switch>
          <Route exact path="/" component={Dashboard} />
          <Route path="/gathering/:id" component={GatheringDetail} />
          <Route path="/login" component={Login} />
          <Route path="/register" component={Register} />
        </Switch>
      </div>
    </Router>
  );
};

export default App;