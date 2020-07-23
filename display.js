import $ from 'jquery';

const Display = function(width, height) {
  this.width  = width;
  this.height = height;

  var currentState;

  $('body').css("background-color","gray").css("text-align","center");

  var c = $('<canvas/>').attr("id","canvas").css("image-rendering","pixelated");
  $(document.body).append(c);

  var buffer  = document.createElement("canvas").getContext("2d")
  var context = document.getElementById('canvas').getContext("2d")

  buffer.canvas.width = width
  buffer.canvas.height = height

  setInterval(render, 1000/30);

  setInterval(renderScene, 1000);

  function renderScene() {
    console.log("renderScene")
    buffer.fillStyle = "#101010";
    buffer.fillRect(0, 0, buffer.canvas.width, buffer.canvas.height);
    buffer.strokeStyle = "white"
    buffer.lineWidth = 1
    buffer.strokeRect(10,10,width-20,height-20);
  };

  function render() {
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

}

export {Display}
