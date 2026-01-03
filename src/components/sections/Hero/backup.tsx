export default function Hero() {
  return (
    <section className="min-h-screen flex flex-col justify-center items-center text-center px-4">  
        <h1 className="text-5xl md:text-7xl font-bold mb-4 text-white"> 
            Welcome to My Portfolio
        </h1>
        <p className="text-lg md:text-2xl text-gray-300 max-w-2xl mb-8">   
            I'm Aries Dimasy, a passionate developer specializing in creating beautiful and functional web applications.
        </p>
        <a 
            href="#projects"
            className="bg-blue-600 text-white px-6 py-3 rounded-full text-lg hover:bg-blue-700 transition"
        >
            View My Work         
        </a>
    </section>
  )
}