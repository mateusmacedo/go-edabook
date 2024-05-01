# Notes

## 1. Introduction to Event-Driven Architectures

### Patterns

- Event notifications
- Event-carried state transfer
- Event sourcing

#### Event notification

- minimum state
- identifier **ID**
- exact time of the occurrence

#### Event-carried state transfer

- push model transfer
- identifier **ID**
- exact time of the occurrence
- payload with the data changes

#### Event sourcing

- streams of events
- identifier **ID**
- payload with the data
- exact time of the occurrence
- recreation of the final state

### Core Components

- Event
- Queue
- Producer
- Consumer

#### Event

- it is an occurrence that has happened in the application
- is in the past and it is an immutable fact
- simple value objects that contain state
- equal to another if all the attributes are the same

#### Queue

referred to by a variety of terms, including bus, channel, stream, topic, and others

##### Message queue

- message queue is its lack of event retention
- events put into a message queue have a limited lifetime
- have been consumed or have expired, they are discarded
- useful for simple publisher/subscriber (pub/sub)
- scenarios when the subscribers are actively running or can retrieve the events quickly

##### Event streams

- When you add event retention to a message queue
- may now read event streams starting with the earliest event
- can begin consuming new events as they are added

##### Event stores

- append-only repository for events
- provide optimistic concurrency controls
- not used for message communication
- used in conjunction with event sourcing to track changes to entities

##### Producers

- publish an event representing the change into the appropriate queue
- may include additional metadata along with the event
- metadata is useful for tracking, performance, or monitoring
- publish it without knowing what the consumers might be listening to

##### Consumers

- subscribe to and read events from queues
- organized into groups to share the loadorganized into groups to share the load
- can be individuals reading all events as they are published
- reading from streams may choose to read from the beginning

### Organizing App

#### Application Services

- use events to communicate new states
- triggers and notifications
- publish them and subscribe to them

#### API gateway services

- RESTful API
- Websocket
- gRPC
- implemented as a BFF pattern

#### Clients

- Single Page Applicatiion
- Mobile Application
- Api Integration

#### Note about hexagons

The services in Figure have some combinations of synchronous and asynchronous communication or connections, and all are drawn as hexagons, as depicted in the following diagram:

![Hexagonal representation of a service](../media/hexagon.png)

In a P2P connection as shown in the following diagram, the calling component, Orders, is dependent on the called component, Depot using a `Client` to make communication

![P2P communication](../media/p2p.png)

### Challenges

some challenges that must be overcome for the application to succeed.

#### Eventual consistency

- Changes in the application state may not be immediately available
- Queries may produce stale results until the change has been fully recorded
- An asynchronous application might have to deal with eventual consistency

#### Dual writes

- refers to any time you’re changing the application state in two or more places during an operation
- making a change locally to a database, and then we’re publishing an event either about the event or the event itself
- solution that will have us publish our events into the database alongside the rest of the changes to keep the state change atomic

![Outbox pattern](../media/outbox.png)

#### Distributed and asynchronous workflows

- performing complex workflows across components using events
- making the workflow entirely asynchronous
- component may not have the final state of the application when queried

#### UX

- difficult to return a final result to the user
- Solutions include but are not limited to:
  - polling on the client
  - delivering the result asynchronously using WebSockets
  - creating the expectation the user should check later for the result

#### Component collaboration

![Component collaboratiion](../media/colaboration.png)

- **Choreography:** The components each individually know about the work they must do, and
which step comes next
- **Orchestration:** The components know very little about their role and are called on to do their
part by a centralized orchestrator

#### Debuggability

In Synchronous communication or P2P involves a caller and callee, always knowing what was called and what made the call

- publish an event and not necessarily knowing if anything is consuming
- challenge in tracing an operation across the application components.
