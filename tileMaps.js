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

        var layers = scenes[scene].scene.layers.reduce(function(map, layer) {
          if(layer.compression == 'gzip') {

            var strData     = atob(layer.data);
            var charData    = strData.split('').map(function(x){return x.charCodeAt(0);});
            var binData     = new Uint8Array(charData);
            var strData     = new Uint32Array((pako.inflate(binData)).buffer);

            //var strData     = String.fromCharCode.apply(null, new Uint16Array(data));

            console.log(strData);
            layer['decompressedData']=strData;
            map[layer.name] = layer;
            console.log("layer:"+layer.name+" decompressed to 32attay:"+layer['decompressedData'])
          } else {
            console.log("unknown compression format '"+layer.compression+"' for layer "+layer.name+" in "+scene+" not found")
          }
          return map;
        }, {});
        if (layer in layers) {

        } else {
          console.log("layer "+layer+" in "+scene+" not found")
        }
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
