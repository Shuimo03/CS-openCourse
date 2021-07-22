import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.Assertions.*;
import java.util.HashSet;

public class LRUTest {

  @Test
  public void lruTest(){
    LRU lru = new LRU(3);
    lru.refer(7);
    lru.refer(0);
    lru.refer(1);
    lru.refer(2);
    lru.refer(0);
    lru.refer(3);
    lru.refer(0);
    lru.refer(4);
  }
}
