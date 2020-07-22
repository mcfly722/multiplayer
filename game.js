import {Controller} from './controller.js';
import {Display} from './display.js';

var jwtDecode = require('jwt-decode');

function Game() {

  var playerId;
  var controller;
  var display;
  var token;

  var state;

  $.ajax({
    url: "api/login",
  }).done(function(data) {
    token = data;
    playerId = jwtDecode(token).Id;
    controller = new Controller(token);
    display = new Display();
    setInterval(update, 1000);
  });


  function update() {
    $.ajax({
      type: "GET",
      url: "api/state",
      headers: { Token: token }
    }).done(function(data) {
      this.state = JSON.parse(data)
      console.log(data)
    });
  }

};

Game.prototype = {
  constructor : Game
};

export {Game}
