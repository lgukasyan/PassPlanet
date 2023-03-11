import Navbar from "./components/Navbar";
import { FcLockLandscape, FcFaq } from 'react-icons/fc';

function App() {
  return (
    <>
      <Navbar />
      <div className="py-32 w-[700px] text-center mx-auto">
        <h2 className="font-extrabold text-5xl leading-17">Start <a class="underline decoration-[#5A4AE3]">protecting</a> <a class="underline decoration-sky-500">your</a> <a class="underline decoration-[#5A4AE3]">passwords</a> with PassPlanet!</h2>
        <div className="text-center text-gray-700 w-[500px] mx-auto">
          <div className="flex justify-center items-center space-x-3">
            <FcLockLandscape size={"50px"} />
            <FcFaq size={"50px"} />
          </div>
          <p><b>PassPlanet</b> has been created for <b>everyone</b>, enjoy protecting your passwords with this software.</p>
          <div className="flex justify-center items-center py-5">
            <a href="#" className="py-2 px-7 bg-[#5A4AE3] hover:bg-gray-500 rounded shadow font-bold text-white">Get Started</a>
          </div>
        </div>
      </div>
    </>
  );
}

export default App;
