package vero

//type Hash struct {
//	ID   int
//	Hash string
//}
//
////func (h Hash)
//
//func read(q int) []Hash {
//	file, err := os.Open("hash.csv")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer func(file *os.File) {
//		err = file.Close()
//		if err != nil {
//			log.Fatal(err)
//		}
//	}(file)
//	content, err := csv.NewReader(file).ReadAll()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	var Hashes []Hash
//
//	for i := 0; i <= q && i < len(content); i++ {
//		fmt.Println(content[len(content)-i-1][0])
//	}
//
//	return Hashes
//}
//
//func generate(q int) {
//	file, err := os.Create("hash.csv")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer func(file *os.File) {
//		err = file.Close()
//		if err != nil {
//			log.Fatal(err)
//		}
//	}(file)
//	writer := csv.NewWriter(file)
//	defer writer.Flush()
//
//	secret := hash("SECRET")
//	for i := 1; i <= q; i++ {
//		err = writer.Write([]string{secret, strconv.Itoa(i)})
//		if err != nil {
//			log.Fatal(err)
//		}
//		secret = hash(secret)
//	}
//}
//
//func hash(s string) string {
//	h := sha256.New()
//	h.Write([]byte(s))
//	bs := h.Sum(nil)
//	return fmt.Sprintf("%x", bs)
//}
