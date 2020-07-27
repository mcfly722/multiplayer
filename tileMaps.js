import $ from 'jquery';
import {Images} from './images.js';

var pako = require('pako');

const TileMaps = function() {

  var images = new Images();

  var scenes = {};

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
                console.log("unknown compression format '"+l.compression+"' for layer "+l.name+" in "+scene+" not found")
              }
            }

            if (l.hasOwnProperty("decompressedData")){

              var sx1 = Math.floor(sx/s.tilewidth)
              var sy1 = Math.floor(sy/s.tileheight)
              var sx2 = Math.floor((sx+w)/s.tilewidth)
              var sy2 = Math.floor((sy+h)/s.tileheight)

              //console.log("put sx1="+sx1+" sy1="+sy1+" sx2="+sx2+" sy2="+sy2);

              for(var ix = sx1;ix<=sx2;ix++){
                for(var iy = sy1;iy<=sy2;iy++){

                  var tile_n = (l['decompressedData'])[iy*l.width+ix]-1
                  var tile_ix = tile_n % tileSet.columns
                  var tile_iy = Math.floor(tile_n/tileSet.columns)


                  var tile_ax1 = sx - sx1*s.tilewidth
                  var tile_ay1 = sy - sy1*s.tileheight



                  //function(ctx,imageSrc,sx,sy,dx,dy,w,h)


                  images.putImage(ctx,tileSet.image,
                    tile_ix * s.tilewidth,
                    tile_iy * s.tileheight,
                    dx+(ix-sx1)*s.tilewidth-tile_ax1,
                    dy+(iy-sy1)*s.tileheight-tile_ay1,
                    s.tilewidth,
                    s.tileheight)

                    //images.putImage(ctx,'tileset1.png',0,0,0,0,200,200)
                    //console.log(" ix="+ix+" iy="+iy+" tile_n="+tile_n+ " tile_ix="+tile_ix+" tile_iy="+tile_iy+" tile_ax1="+tile_ax1+" tile_ay1="+tile_ay1+"  tileSet="+tileSet.image + " putted");

                }
              }






            }
            //console.log("put layer:"+l.name)

          }
        });
      }
    } else {
      $.ajax({
        url: scene,
      }).done(function(data) {
        scenes[scene].scene = data;
        scenes[scene].ready = true;
      });
      scenes[scene] = {ready:false}
    }
  }

}

TileMaps.prototype = {
  constructor : TileMaps
};

export {TileMaps}
