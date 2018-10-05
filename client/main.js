const rp = require('request-promise');

rp.get('http://localhost:8080/players/')
.then(players => {
    console.log("List all players:", players);
});

rp.get('http://localhost:8080/players/1/')
.then(player => {
    console.log("Player 1:", player);
});