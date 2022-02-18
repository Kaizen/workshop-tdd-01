import org.junit.jupiter.api.Assertions.*
import org.junit.jupiter.api.Test

internal class HelloWorldTest {
    
    @Test
    fun instantiates() {
       val h = HelloWorld()
       assertNotNull(h) 
    }
}