export default function Navbar() {
  return (
    <nav>
      <div className="max-w-7xl mx-auto px-4">
        <div className="flex justify-between items-center">
          <div className="flex space-x-4 items-center">
            <div className="mr-1">
              <a href="#" className="flex py-5 px-3">
                <img src={"/images/PassPlanet.png"} width="250px" />
              </a>
            </div>
            <div className="flex items-center space-x-3 hidden md:flex">
              <a href="#" className="py-4 px-3 font-bold hover:text-gray-600">Home</a>
              <a href="#" className="py-3 px-3 font-bold hover:text-gray-600">About</a>
            </div>
          </div>
          <div>
            <div className="hidden md:flex">
              <a href="#" className="py-3 px-3 font-bold hover:text-gray-600">Login</a>
              <a href="#" className="py-3 px-3 bg-[#5A4AE3] hover:bg-gray-500 rounded shadow font-bold text-white">Sign Up</a>
            </div>
          </div>
        </div>
      </div>
    </nav>
  );
}