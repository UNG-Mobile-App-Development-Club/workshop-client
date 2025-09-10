use workshop_client::TriviaResponse;

fn main() {
    let url = "https://opentdb.com/api.php?amount=10";
    let res = reqwest::blocking::get(url)
        .expect("Failed to get trivia response");
    let data: TriviaResponse = res.json().expect("Failed to parse response");

    println!("{:#?}", data);
}
