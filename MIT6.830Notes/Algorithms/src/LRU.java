import java.util.Deque;
import java.util.HashSet;
import java.util.Iterator;
import java.util.LinkedList;

public class LRU {
    //store keys of cache
    private Deque<Integer> deque;
    // store references of key in cache
    private HashSet<Integer> hashSet;

    // maximum capacity of cache
    private final int CACHE_SIZE;

    LRU(int capacity){
        deque = new LinkedList<>();
        hashSet = new HashSet<>();
        CACHE_SIZE = capacity;
    }

    /* Refer the page within the LRU cache */
    public void refer(int  page){
        if(!hashSet.contains(page)){
            if(deque.size() == CACHE_SIZE){
                int last = deque.removeLast();
                hashSet.remove(last);
            }
        }
        else{
            /* The found page may not be always the last element, even if it's an
               intermediate element that needs to be removed and added to the start
               of the Queue */
            deque.remove(page);
        }
        deque.push(page);
        hashSet.add(page);
    }

    public void display(){
        Iterator<Integer> iterator = deque.iterator();
        while (iterator.hasNext()){
            System.out.print(iterator.next()+" ");
        }
    }
}
