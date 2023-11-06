import java.util.Arrays;
import java.util.Base64;
import java.util.Random;

import javax.crypto.Cipher;
import javax.crypto.spec.IvParameterSpec;
import javax.crypto.spec.SecretKeySpec;
 
public class Main {
    public static void main(String[] args) throws Exception {
        final var keyStr = "0123456789abcdef0123456789abcdef";

        if (args.length == 0) {
            final var plain = "Hello";
            final var encrypted = encrypt(plain, keyStr);
            final var decryted = decrypt(encrypted, keyStr);
            System.out.printf("Plain: %s\nEncrypted: %s\nDecrypted: %s\n", plain, encrypted, decryted);
            return;
        }

        if (args.length != 2) {
            System.out.println("java Main.java <encrypt | decrypt> \"content\"");
            return;
        }

        final var kind = args[0];

        if ("encrypt".equals(kind)) {
            System.out.println(encrypt(args[1], keyStr));
            return;
        }

        if ("decrypt".equals(kind)) {
            System.out.println(decrypt(args[1], keyStr));
            return;
        }

        System.out.println("java Main.java <encrypt | decrypt> \"content\"");
    }

    public static String encrypt(String data, String keyStr) throws Exception {
        //final var keyInHexBytes = HexFormat.of().parseHex(key);
        final var key = keyStr.getBytes("UTF-8");
        final var initVector = generateIv();
        final var ivSpec = new IvParameterSpec(initVector.getBytes("UTF-8"));
        final var keySpec = new SecretKeySpec(key, "AES");
        final var ciper = Cipher.getInstance("AES/CBC/PKCS5PADDING");
        ciper.init(Cipher.ENCRYPT_MODE, keySpec, ivSpec);

        final var encryptedBytes = ciper.doFinal(data.toString().getBytes());
        final var finalArray = new byte[initVector.length() + encryptedBytes.length];
        
        System.arraycopy(initVector.getBytes(), 0, finalArray, 0, initVector.getBytes().length);
        System.arraycopy(encryptedBytes, 0, finalArray, initVector.getBytes().length, encryptedBytes.length);
        
        final var encrypted = Base64.getEncoder().encodeToString(finalArray);
        return encrypted;
    }

    public static String decrypt(String encrypted, String keyStr) throws Exception {
        //final var key = HexFormat.of().parseHex(keyStr);
        final var key = keyStr.getBytes("UTF-8");
        final var keySpec = new SecretKeySpec(key, "AES");

        final var encryptedBytes = Base64.getDecoder().decode(encrypted);
        final var initVector = Arrays.copyOfRange(encryptedBytes, 0, 16);
        final var encryptedPayload = Arrays.copyOfRange(encryptedBytes, initVector.length, encryptedBytes.length);

        final var ivSpec = new IvParameterSpec(initVector);
        final var ciper = Cipher.getInstance("AES/CBC/PKCS5PADDING");
        ciper.init(Cipher.DECRYPT_MODE, keySpec, ivSpec);

        final var decryptedBytes = ciper.doFinal(encryptedPayload);
        final var decrypted = new String(decryptedBytes);

        return decrypted;
    }
    
    public static String generateIv() {
        final int lowAsciiLimit = 47; // '/'
        final int highAsciiLimit = 126; // 'z'
        
        final int ivLength = 16;
        final var random = new Random();
        final var ivBuffer = new StringBuilder(ivLength);
        
        for (int i=0; i<ivLength; i++) {
            final int randomNumber = lowAsciiLimit + (int) (random.nextFloat() * (highAsciiLimit - lowAsciiLimit));
            ivBuffer.append((char) randomNumber);
        }
        return ivBuffer.toString();
    }

    /*public static String decode(String base64Text, byte[] key) throws Exception {
        byte[] inputArr = Base64.getDecoder().decode(base64Text);
        SecretKeySpec skSpec = new SecretKeySpec(key, "AES");
        Cipher cipher = Cipher.getInstance("AES/CBC/PKCS5PADDING");
        int blockSize = cipher.getBlockSize();
        IvParameterSpec iv = new IvParameterSpec(Arrays.copyOf(inputArr, blockSize));
        byte[] dataToDecrypt = Arrays.copyOfRange(inputArr, blockSize, inputArr.length);
        cipher.init(Cipher.DECRYPT_MODE, skSpec, iv);
        byte[] result = cipher.doFinal(dataToDecrypt);
        return new String(result, StandardCharsets.UTF_8);
    }*/
}

