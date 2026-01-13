let seeBio = false

function viewBiographie() {
    const bioScreen = document.getElementById("bioScreen").style.display
    if ( seeBio === false ) {
        document.getElementById("bioScreen").style.display = "flex"
        seeBio = true
        document.getElementById("btnSeeBio").textContent = "Cacher la biographie"
    } else {
        document.getElementById("bioScreen").style.display = "none"
        seeBio = false            
    }
}

function viewBiographie() {
    document.getElementById("bioOverlay").style.display = "block";
    document.getElementById("bioScreen").style.display = "block";
    seeBio = true;
}

function closeBio() {
    document.getElementById("bioOverlay").style.display = "none";
    document.getElementById("bioScreen").style.display = "none";
    seeBio = false;
}