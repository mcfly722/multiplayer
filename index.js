import $ from 'jquery';
var jwtDecode = require('jwt-decode');

$.ajax({
  url: "api/login",
}).done(function(data) {
  token = data;
  playerId = jwtDecode(data).Id;
  startGame()
});


var actualKeys = {
  'ArrowUp':false,
  'ArrowDown':false,
  'ArrowLeft':false,
  'ArrowRight':false,
  'Space':false
}

var serverKeys;

function updateMovements(){
  if (JSON.stringify(serverKeys) != JSON.stringify(actualKeys)) {

    serverKeys = JSON.parse(JSON.stringify(actualKeys));

    $.ajax({
      type: "POST",
      url: "api/movement",
      headers: { Token: token },
      data: JSON.stringify(serverKeys)
    });

    console.log("updated to "+JSON.stringify(serverKeys));
  }
}

var playerId;
var token;

function startGame() {

  window.addEventListener('keyup', function(event) {
    if (event.code in actualKeys) {
      actualKeys[event.code] = false
    }
  }, false);

  window.addEventListener('keydown', function(event) {
    if (event.code in actualKeys) {
      actualKeys[event.code] = true
    }
  }, false);

  setInterval(updateMovements, 100);
}
