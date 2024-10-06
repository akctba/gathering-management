import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';

function App() {
  return (
    <Router>
      <div className="App">
        <header className="App-header">
          <h1>Gathering Management</h1>
        </header>
        <main>
          <Switch>
            <Route exact path="/" component={Home} />
            {/* Add more routes here */}
          </Switch>
        </main>
      </div>
    </Router>
  );
}

function Home() {
  return <h2>Welcome to Gathering Management</h2>;
}

export default App;