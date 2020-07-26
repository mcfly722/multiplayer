import $ from 'jquery';
import {TileMaps} from './tileMaps.js';
import {Images} from './images.js';

const Display = function(width, height) {

  var images = new Images();
  var tileMaps = new TileMaps();

  var playerId,currentWorldState;

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


  function renderHero() {
    if (playerId !== undefined) {
      //console.log("world:"+JSON.stringify(currentWorldState))
      //console.log("playerID:"+playerId)
      var spriteSetNum = currentWorldState.Players[playerId].SpriteSetNum;
      images.putImage(buffer,'player'+spriteSetNum+'.png',0,0,buffer.canvas.width/2-16,buffer.canvas.height/2,32,32)
    }
  }

  var a=0;

  function renderScene() {
    buffer.fillStyle = "#101010";
    buffer.fillRect(0, 0, buffer.canvas.width, buffer.canvas.height);

    if (currentWorldState !== undefined){
      a++;
      var xx = Math.round(40+40*Math.sin(a/30))
      var yy = Math.round(40+40*Math.cos(a/30))

      tileMaps.putLayer(buffer,currentWorldState.SceneName,"Layer1",0,0,xx,yy,200,200);
      tileMaps.putLayer(buffer,currentWorldState.SceneName,"Layer2",0,0,xx,yy,200,200);
    }

    buffer.strokeStyle = "white"
    buffer.lineWidth = 1
    buffer.strokeRect(10,10,width-20,height-20);

  };

  function render() {
    renderScene();
    renderHero();

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
    console.log("apply new state")
    playerId = playerId_;
    currentWorldState = worldState_;
  }

}

export {Display}
