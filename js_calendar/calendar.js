/*
    カレンダープログラム

    使用方法：
        node calendar.js [-m month]
    使用例：
        node calendar.js
        node calendar.js -m 6
*/

const args = process.argv.slice(3);  // 引数取得

showCallender(args);

function showCallender(month) {

    if (month.length == 0) {
        // 引数に月指定なしの場合、現在日付（月）とする
        month = new Date().getMonth() + 1;  // monthは0から始まるので、+1する
    }

    if (!(1 <= month && month <= 12)) {
        // 引数が不正な月を指定している場合、エラー
        console.log(`${month} is neither a month number (1..12) nor a name`);
        process.exit();
    }

    const monthStr = String(month).replace(/0/g, '');
    const year = new Date().getFullYear();
    const date = new Date(year, monthStr - 1, 1);
    const farstDate = date.getDate();
    const lastDate = new Date(year, monthStr, 0).getDate();  // 0を指定で、最終日取得
    const DayOfWeek = date.getDay();

    let calendar = '';

    // 最初の日まで空白を入力
    for (let i = 0; i < DayOfWeek; i++) {
        calendar += '    ';
    }

    // 月の日数を入力
    for (let i = farstDate; i <= lastDate; i++) {
        calendar += i.toString().padStart('2', ' ') + '  ';

        let currentDayOfWeek = new Date(year, month - 1, i).getDay();
        if (currentDayOfWeek == 6) {
            // 土曜日になったら、改行する
            calendar += '\n';
        }
    }

    console.log(`        ${monthStr}月  ${year}`);
    console.log('日  月  火  水  木  金  土');
    console.log(calendar);
}