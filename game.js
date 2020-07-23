import {Controller} from './controller.js';
import {Display} from './display.js';

var jwtDecode = require('jwt-decode');

function Game() {

  var playerId;
  var controller;
  var display;
  var token;

  var reloginRequired = true;

  function login(){
    $.ajax({
      url: "api/login",
    }).done(function(data) {
      token = data;
      playerId = jwtDecode(token).Id;
      controller = new Controller(token);
      display = new Display(480, 300);
      display.resize();
      setInterval(update, 1000);
      reloginRequired = false;
    });
  }

  setInterval(checkThatLogined, 1000);

  function checkThatLogined(){
      if (reloginRequired){
        login();
      }
  }

  function update() {
    if(!reloginRequired) {
      $.ajax({
        type: "GET",
        url: "api/state",
        headers: { Token: token }
      }).done(function(data) {
        display.applyNewState(playerId, JSON.parse(data));
      }).fail(function(jqXHR){
        if(jqXHR.status == 401) {
          reloginRequired = true;
        }
      });
    }
  }
};

Game.prototype = {
  constructor : Game
};

export {Game}
