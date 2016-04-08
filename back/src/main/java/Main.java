import net.codestory.http.WebServer;
import net.codestory.http.filters.log.LogRequestFilter;

import java.net.InetAddress;
import java.net.UnknownHostException;

public class Main {
    public static void main(String[] args) throws UnknownHostException {
        String hostname = InetAddress.getLocalHost().getHostName();

        new WebServer().configure(routes -> routes
                .filter(LogRequestFilter.class)
                .get("/", new WordResponse("TEST", hostname))
        ).start();
    }

    public static class WordResponse {
        String word;
        String hostname;

        WordResponse(String word, String hostname) {
            this.word = word;
            this.hostname = hostname;
        }
    }
}
