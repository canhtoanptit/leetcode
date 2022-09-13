var twoSum = function(nums, target) {
    const hashmap = new Map();
    let output;
    
    for (let i = 0; i < nums.length; i++) {
        const num = nums[i];
        const value = hashmap.get(target - num)
        if(hashmap.has(target - num)) {
            output = [value, i]
        } else {
            hashmap.set(num, i);
        }
    }
    
    return output;
};
