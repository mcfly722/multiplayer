import $ from 'jquery';
import {Images} from './images.js';

var pako = require('pako');

const TileMaps = function() {

  var images = new Images();

  var scenes = {};

  this.getScene = function(sceneName){
    if (sceneName in scenes) {
      if(scenes[sceneName].ready) {
        return scenes[sceneName].scene
      }
    }
    return null
  }

  this.loadScene = function(sceneName){
    if (!(sceneName in scenes)){
      $.ajax({
        url: sceneName
      }).done(function(data) {
        scenes[sceneName].scene = data;
        scenes[sceneName].ready = true;
      });
      scenes[sceneName] = {ready:false}
      console.log("start to load "+sceneName)
    }
  }

  this.putLayer = function(ctx,scene,layerName,sx,sy,dx,dy,w,h){
    if (scene in scenes) {
      if(scenes[scene].ready) {
        var tileSet = scenes[scene].scene.tilesets[0];
        var s = scenes[scene].scene
        s.layers.forEach(function(l){
          if (l.name == layerName) {

            // decompress if not decompressed yet
            if (!l.hasOwnProperty("decompressedData")){
              if(l.compression == 'gzip') {

                var strData     = atob(l.data);
                var charData    = strData.split('').map(function(x){return x.charCodeAt(0);});
                var binData     = new Uint8Array(charData);
                var strData     = new Uint32Array((pako.inflate(binData)).buffer);

                l['decompressedData']=strData;
                console.log("layer:"+l.name+" decompressed")
              } else {
                if(l.compression !== undefined){
                  console.log("unknown compression format '"+l.compression+"' for layer "+l.name+" in "+scene+" not found")
                }
              }
            }

            if (l.hasOwnProperty("decompressedData")){
              var sx1 = Math.floor(sx/s.tilewidth)
              var sy1 = Math.floor(sy/s.tileheight)
              var sx2 = Math.floor((sx+w)/s.tilewidth)
              var sy2 = Math.floor((sy+h)/s.tileheight)
              for(var ix = sx1;ix<=sx2;ix++){
                for(var iy = sy1;iy<=sy2;iy++){
                  var tile_n = (l['decompressedData'])[iy*l.width+ix]-1
                  var tile_ix = tile_n % tileSet.columns
                  var tile_iy = Math.floor(tile_n/tileSet.columns)
                  var tile_ax1 = sx - sx1*s.tilewidth
                  var tile_ay1 = sy - sy1*s.tileheight
                  images.putImage(ctx,tileSet.image,
                    tile_ix * s.tilewidth,
                    tile_iy * s.tileheight,
                    dx+(ix-sx1)*s.tilewidth-tile_ax1,
                    dy+(iy-sy1)*s.tileheight-tile_ay1,
                    s.tilewidth,
                    s.tileheight)
                }
              }
            }

            if(l.objects !== undefined){
              l.objects.forEach(obj => {
                if(obj.polygon !== undefined){

                  ctx.beginPath();
                  ctx.moveTo(obj.x+obj.polygon[0].x-sx, obj.y+obj.polygon[0].y-sy);

                  obj.polygon.forEach(line=>{
                    ctx.lineTo(obj.x+line.x-sx, obj.y+line.y-sy);
                  })

                  ctx.lineTo(obj.x-sx, obj.y-sy);
                  ctx.strokeStyle = "red"
                  ctx.lineWidth = 2
                  ctx.stroke();

                } else {
                  ctx.strokeStyle = "blue"
                  ctx.lineWidth = 2
                  ctx.strokeRect(obj.x-sx,obj.y-sy,obj.width,obj.height);
                }
              });
            }
          }
        });
      }
    } else {
      this.loadScene(scene);
    }
  }

}

TileMaps.prototype = {
  constructor : TileMaps
};

export {TileMaps}
