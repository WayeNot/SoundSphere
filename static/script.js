history.replaceState("", "", "SoundSphère");
let seeBio = false
function viewBiographie() {
    alert("iefhef")
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