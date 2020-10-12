import $ from 'jquery';
import {TileMaps} from './tileMaps.js';
import {Images} from './images.js';

const Display = function(width, height) {

  var camera;

  var images = new Images();
  var tileMaps = new TileMaps();

  var playerId, currentWorldState;

  this.width  = width;
  this.height = height;

  $('body').css("background-color","gray").css("text-align","center");

  var c = $('<canvas/>').attr("id","canvas").css("image-rendering","pixelated");
  $(document.body).append(c);

  var buffer  = document.createElement("canvas").getContext("2d")
  var context = document.getElementById('canvas').getContext("2d")

  buffer.canvas.width = width
  buffer.canvas.height = height

  setInterval(render, 1000/60);

  setInterval(renderScene, 1000);


  function renderPlayers() {

    for (var playerId in currentWorldState.Players){

      var player = currentWorldState.Players[playerId]
      var spriteSetNum = player.SpriteSetNum;

      var pos = {x:Math.round(player.X-camera.x+buffer.canvas.width/2),y:Math.round(player.Y-camera.y+buffer.canvas.height/2)}
      images.putImage(buffer,'player'+spriteSetNum+'.png',
        0,0,
        pos.x - 16,pos.y -28,
        32,32)

        buffer.beginPath();
        buffer.strokeStyle = "red"
        //buffer.fillStyle = "red"
        buffer.lineWidth = 2
        buffer.arc(
          pos.x,pos.y,
          5, 0, 2 * Math.PI);
        buffer.stroke();
    }
  }

  var a=0;

  function renderScene() {
    buffer.fillStyle = "#101010";
    buffer.fillRect(0, 0, buffer.canvas.width, buffer.canvas.height);

    tileMaps.loadScene(currentWorldState.SceneName);

    var scene = tileMaps.getScene(currentWorldState.SceneName);
    if (scene != null) {
      scene.layers.
      //filter(layer => layer.visible == true).
      forEach(layer => {

        tileMaps.putLayer(
          buffer,
          currentWorldState.SceneName,
          layer.name,
          Math.round(camera.x-buffer.canvas.width/2),Math.round(camera.y-buffer.canvas.height/2),
          0,0,
          buffer.canvas.width,buffer.canvas.height);

          if (layer.name == 'Main') {
            renderPlayers();
          }

      });
    }


    buffer.strokeStyle = "white"
    buffer.lineWidth = 1
    buffer.strokeRect(10,10,width-20,height-20);

  };

  function updateCamera(){
    if (camera !== undefined){
      var player = currentWorldState.Players[playerId]
      camera.x = camera.x+(player.X-camera.x)*0.1
      camera.y = camera.y+(player.Y-camera.y)*0.1
    }
  }


  function render() {
    if (currentWorldState !== undefined && camera !== undefined && playerId !== undefined){
      updateCamera();
      renderScene();
    }

    context.drawImage(buffer.canvas, 0, 0, buffer.canvas.width, buffer.canvas.height, 0, 0, context.canvas.width, context.canvas.height);
  };

  this.resize = function(event) {
      resizeGameWindow(document.documentElement.clientWidth - 32, document.documentElement.clientHeight - 32, height / width);
      render();
  };

  function resizeGameWindow(width, height, height_width_ratio) {
    if (height / width > height_width_ratio) {
        context.canvas.height = width * height_width_ratio;
        context.canvas.width = width;
      } else {
        context.canvas.height = height;
        context.canvas.width = height / height_width_ratio;
      }
      context.scale(1,1)
      context.imageSmoothingEnabled = false;
    };

  this.handleResize = (event) => { this.resize(event); };
  window.addEventListener("resize",  this.handleResize);


  this.applyNewState = function(playerId_, worldState_) {
    //console.log("apply new state")
    playerId = playerId_;
    currentWorldState = worldState_;

    if (camera === undefined){
      camera = {'x':currentWorldState.Players[playerId].X, 'y':currentWorldState.Players[playerId].Y}
    }

  }

}

export {Display}
