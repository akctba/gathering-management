import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';

const Dashboard = () => {
  const [gatherings, setGatherings] = useState([]);

  useEffect(() => {
    // TODO: Fetch gatherings from the backend API
    // For now, we'll use dummy data
    setGatherings([
      { id: 1, name: 'Summer Barbecue', date: '2024-07-15' },
      { id: 2, name: 'Birthday Party', date: '2024-08-20' },
    ]);
  }, []);

  return (
    <div className="dashboard">
      <h1>My Gatherings</h1>
      <ul>
        {gatherings.map(gathering => (
          <li key={gathering.id}>
            <Link to={`/gathering/${gathering.id}`}>
              {gathering.name} - {gathering.date}
            </Link>
          </li>
        ))}
      </ul>
      <button>Create New Gathering</button>
    </div>
  );
};

export default Dashboard;