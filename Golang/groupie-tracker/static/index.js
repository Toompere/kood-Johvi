function initMap() {
    const map = new google.maps.Map(document.getElementById("map"), {
      zoom: 1,
      center: new google.maps.LatLng(36.780651, 5.328462),
      mapTypeControlOptions: {
        position: google.maps.ControlPosition.RIGHT_BOTTOM,
      },
      zoomControlOptions: {
        position: google.maps.ControlPosition.TOP_RIGHT,
      },
      streetViewControlOptions: {
        position: google.maps.ControlPosition.LEFT_TOP,
      },
      fullscreenControlOptions: {
        position: google.maps.ControlPosition.LEFT_BOTTOM,
      },
      });

    locations.forEach(location => {
        var marker = new google.maps.Marker({
            position: location.position,
            map: map,
            title: location.title
        });
        marker.addListener("click", () => {
            map.setZoom(8);
            map.setCenter(marker.getPosition());
              });
    });
 
    
    }
  
  
  window.initMap = initMap;