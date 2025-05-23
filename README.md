# golang-parking-lot
Code to solve parking lot problem (Vocagame)

# Problem 
I own a parking lot that can hold up to 'n' cars at any given point in time. Each slot is given a number starting at 1
increasing with increasing distance from the entry point in steps of one. I want to create an automated ticketing
system that allows my customers to use my parking lot without human intervention. When a car enters my parking lot,
I want to have a ticket issued to the driver. The ticket issuing process includes us documenting the registration
number (number plate) and the colour of the car and allocating an available parking slot to the car before actually
handing over a ticket to the driver (we assume that our customers are nice enough to always park in the slots
allocated to them). The customer should be allocated a parking slot which is nearest to the entry. At the exit the
customer returns the ticket with the time the car was parked in the lot, which then marks the slot they were using as
being available. Total parking charge should be calculated as per the parking time. Charge applicable is $10 for first 2
hours and $10 for every additional hour. We interact with the system via a simple set of commands which produce a
specific output. 

# Prerequisites
Before running the project, ensure you have the following installed:

1. GO version : go1.21.0 windows/amd64 or above.

# Setup and Installation
Follow these steps to clone and run the project locally:
1. Clone the Repository
Start by cloning the repository to your local machine:
<div class="code-container">
  <pre id="command-text">
    git clone https://github.com/your-username/your-project-name.git
    cd your-project-name
  </pre>
</div>

2. Run the Application Once you have installed dependencies from the terminal using:
<div class="code-container">
  <pre id="command-text">
    go run main.go input_parking_lot.txt
  </pre>
</div>