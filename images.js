const Images = function() {

  var images = {};

  setInterval(loadRequestedImages, 1000);

  function loadRequestedImages() {
    Object.keys(images).forEach(function(key){
      if (key in images) {
        if (images[key].requested == false) {
          images[key].img = new Image();
          images[key].img.onload = function(data) {
              images[key].ready = true;
          };
          images[key].img.src = key;
          images[key].requesed = true;
        }
      }
    })
  }

  this.putImage = function(ctx,imageSrc,sx,sy,dx,dy,w,h){
      if(imageSrc in images){
        if (images[imageSrc].ready){
          ctx.drawImage(images[imageSrc].img, sx,sy,w,h,dx,dy,w,h);
        }
      } else {
        images[imageSrc] = {requested:false, ready:false}
      }
  }

}

Images.prototype = {
  constructor : Images
};

export {Images}
