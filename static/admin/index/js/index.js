function getRandomBrightColor() {
    let brightColors = '0123456789ABCDEF';
    let color = '#';
    // 每个位置随机索引限制在后8位
    for (let i = 0; i < 6; i++) {
        color += brightColors[Math.floor(Math.random() * 8 + 8)];
    }
    return color;
}