import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';

const GatheringDetail = () => {
  const { id } = useParams();
  const [gathering, setGathering] = useState(null);

  useEffect(() => {
    // TODO: Fetch gathering details from the backend API
    // For now, we'll use dummy data
    setGathering({
      id: id,
      name: 'Summer Barbecue',
      description: 'Annual summer gathering with friends',
      date: '2024-07-15',
      location: 'Central Park',
      invitees: [
        { id: 1, name: 'John Doe', email: 'john@example.com', rsvpStatus: 'accepted' },
        { id: 2, name: 'Jane Smith', email: 'jane@example.com', rsvpStatus: 'pending' },
      ],
    });
  }, [id]);

  if (!gathering) {
    return <div>Loading...</div>;
  }

  return (
    <div className="gathering-detail">
      <h1>{gathering.name}</h1>
      <p>{gathering.description}</p>
      <p>Date: {gathering.date}</p>
      <p>Location: {gathering.location}</p>
      
      <h2>Invitees</h2>
      <ul>
        {gathering.invitees.map(invitee => (
          <li key={invitee.id}>
            {invitee.name} ({invitee.email}) - RSVP: {invitee.rsvpStatus}
          </li>
        ))}
      </ul>
      
      <button>Add Invitee</button>
      <button>Manage Food Plates</button>
      <button>Manage Beverages</button>
    </div>
  );
};

export default GatheringDetail;